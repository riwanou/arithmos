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
        keys_data.append(extract_bench_data(data))

    avg_data = keys_data[0]
    for data in keys_data[1:]:
        avg_data['Time'] += data['Time']
    avg_data['Time'] /= 5

    avg_data = avg_data.sort_values('Value')
    avg_data['Name'] = name
    
    return avg_data

def gen_plot(df, name_patterns, col_names, filename, avg=True):
    data_frames = []
    for i in range(0, len(name_patterns)):
        extracted_df = df[df['Name'].str.contains(name_patterns[i])]
        avg_df = []
        if avg:
            avg_df = avg_keys_bench(extracted_df, col_names[i])
        else:
            avg_df = keys_bench(extracted_df, col_names[i])
        data_frames.append(avg_df)

    plt.clf()
    final_df = pd.concat(data_frames)
    ax = sns.lineplot(x='Value', y='Time', hue='Name', data=final_df)
    ax.set(xlabel="nombre de clés", ylabel="Temps (ms)")
    plt.savefig(filename, dpi=300) 

# Ajout
df = pd.read_table("bench_output", header=None, 
                   names=["Name", "Iteration", "Time"])
gen_plot(df, ['AjoutIteratif/heapTree'], ['min heap tree'], 
         'plots/heap_ajout_iteratif')

# Construction
df = pd.read_table("bench_output", header=None, 
                   names=["Name", "Iteration", "Time"])
gen_plot(df, ['Construction/heapBinomial'], ['min heap binomial'], 
         'plots/heap_construction')

# Union
df = pd.read_table("bench_output", header=None, 
                   names=["Name", "Iteration", "Time"])
gen_plot(df, ['Union/heapBinomial'], ['min heap binomial'], 
         'plots/heap_union', avg=False)

