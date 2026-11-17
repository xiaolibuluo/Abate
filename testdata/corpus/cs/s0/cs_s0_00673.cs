using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public delegate void Notifier6280(int value);

    public sealed class Source6280
    {
        public event Notifier6280 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale6281<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public async Task<int> Aggregate6282(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public async Task<int> Aggregate6283(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public static int Scale6284<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public sealed class SlotMap6285
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
}
