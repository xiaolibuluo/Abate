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

    public static int Scale18400<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public static int Scale18401<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public sealed class SlotMap18402
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
    public async Task<int> Aggregate18403(IEnumerable<int> source)
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
    public delegate void Notifier18404(int value);

    public sealed class Source18404
    {
        public event Notifier18404 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
