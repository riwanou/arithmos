import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt

def avg_keys_bench(df, name):
    keys_data = []
    for i in range(1, 6):
        data = df[df['Name'].str.contains('jeu_' + str(i))].copy()
        data['Value'] = data['Name'].str.extract('cles_(\d*)')
        data['Value'] = pd.to_numeric(data['Value'])
        data['Time'] = data['Time'].str.extract('(.*)ns/op')
        data['Time'] = pd.to_numeric(data['Time']) / 1000000
        data = data.reset_index()
        keys_data.append(data)

    avg_data = keys_data[0]
    for data in keys_data[1:]:
        avg_data['Time'] += data['Time']
    avg_data['Time'] /= 5

    avg_data = avg_data.sort_values('Value')
    avg_data['Name'] = name
    
    return avg_data

def gen_plot(df, name_patterns, col_names, filename):
    data_frames = []
    for i in range(0, len(name_patterns)):
        extracted_df = df[df['Name'].str.contains(name_patterns[i])]
        avg_df = avg_keys_bench(extracted_df, col_names[i])
        data_frames.append(avg_df)

    final_df = pd.concat(data_frames)
    ax = sns.lineplot(x='Value', y='Time', hue='Name', data=final_df)
    ax.set(xlabel="nombre de clés", ylabel="Temps (ms)")
    plt.savefig(filename, dpi=300) 

# get data from bench file
df = pd.read_table("bench/bench_bin_tree", header=None, 
                   names=["Name", "Iteration", "Time"])
df = df.drop('Iteration', axis=1)

# plt.savefig("plots/heap_ajout_iteratif") 
gen_plot(df, ['Ajout/heapTree'], ['min heap tree'], 
         'plots/heap_ajout_iteratif')


