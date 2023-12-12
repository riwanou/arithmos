import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt

def extract_bench_data(data):
    data['Value'] = data['Name'].str.extract('cles_(\d*)')
    data['Value'] = pd.to_numeric(data['Value'])
    data['Time'] = data['Time'].str.extract('(.*)ns/op')
    data['Time'] = pd.to_numeric(data['Time']) / 1000000
    data = data.dropna()
    data = data.reset_index()
    return data

def keys_bench(df, name):
    data = df.copy()
    data = extract_bench_data(data)

    data = data.sort_values('Value')
    data['Name'] = name

    return data

def avg_keys_bench(df, name):
    keys_data = []
    for i in range(1, 6):
        data = df[df['Name'].str.contains('jeu_' + str(i))].copy()
        if len(data) > 0:
            keys_data.append(extract_bench_data(data))

    if len(keys_data) == 0:
        return []

    avg_data = keys_data[0]
    for data in keys_data[1:]:
        avg_data['Time'] += data['Time']
    avg_data['Time'] /= 5

    extra_data = df[df['Name'].str.contains('extra_jeu')].copy()
    if len(extra_data) > 0:
        avg_data = pd.concat([extract_bench_data(extra_data), avg_data])

    avg_data = avg_data.sort_values('Value')
    avg_data['Name'] = name

    return avg_data

def get_dataframes(df, name_patterns, col_names, filename, avg=True):
    data_frames = []
    for i in range(0, len(name_patterns)):
        extracted_df = df[df['Name'].str.contains(name_patterns[i])]
        avg_df = []
        if avg:
            avg_df = avg_keys_bench(extracted_df, col_names[i])
        else:
            avg_df = keys_bench(extracted_df, col_names[i])
        if len(avg_df) > 0:
            data_frames.append(avg_df)
    return data_frames

def gen_plot(df, name_patterns, col_names, filename, avg=True):
    data_frames = get_dataframes(df, name_patterns, col_names, filename, avg)
    if len(data_frames) == 0:
        return

    plt.clf()
    final_df = pd.concat(data_frames)
    ax = sns.lineplot(x='Value', y='Time', hue='Name', data=final_df)
    ax.set(xlabel="nombre de clés", ylabel="Temps (ms)")
    plt.savefig(filename, dpi=300) 

def gen_bar_plot(df, name_patterns, col_names, filename, avg=True):
    data_frames = get_dataframes(df, name_patterns, col_names, filename, avg)
    if len(data_frames) == 0:
        return

    plt.clf()
    final_df = pd.concat(data_frames)
    ax = sns.barplot(x="Value", y="Time", hue='Name', data=final_df, gap=0.05)
    ax.set(xlabel="nombre de clés", ylabel="Temps (ms)")
    plt.savefig(filename, dpi=300) 


df = pd.read_table("bench_output", header=None, 
                   names=["Name", "Iteration", "Time"])

# Heaps ajout tree
# gen_plot(df, 
#          ['AjoutIteratif/heapTree', 'Construction/heapTree'], 
#          ['ajout min heap tree', 'construction min heap tree'], 
#          'plots/ajout_tree')

# # Heaps ajout array
gen_plot(df, 
         ['AjoutIteratif/heapArray', 'Construction/heapArray'], 
         ['ajout min heap array', 'construction min heap array'], 
         'plots/ajout_array')

# # Heaps Construction
gen_plot(df, 
         ['Construction/heapBinomial', 'Construction/heapTree', 'Construction/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/construction')

# Heaps Union
gen_plot(df, 
         ['Union/heapBinomial', 'Union/heapTree', 'Union/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/heap_union', avg=False)

# Heap Binomial Union
gen_plot(df, 
         ['Union/heapBinomial'], 
         ['min heap binomial'], 
         'plots/heap_binomial_union', avg=False)

# Shakespeare

# Ajout
gen_bar_plot(df, 
         ['AjoutWords/heapBinomial', 'AjoutWords/heapTree', 'AjoutWords/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/words_ajout', avg=False)

# Construction
gen_bar_plot(df, 
         ['ConstructionWords/heapBinomial', 'ConstructionWords/heapTree', 
             'ConstructionWords/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/words_construction', avg=False)

# SupprMin
gen_bar_plot(df, 
         ['SupprMinWords/heapBinomial', 'SupprMinWords/heapTree', 
             'SupprMinWords/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/words_supprmin', avg=False)

# SupprMin
gen_bar_plot(df, 
         ['UnionWords/heapBinomial', 'UnionWords/heapTree', 
             'UnionWords/heapArray'], 
         ['min heap binomial', 'min heap tree', 'min heap array'], 
         'plots/words_union', avg=False)

