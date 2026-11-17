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

    public sealed class SlotMap8700
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
    public async Task<int> Aggregate8701(IEnumerable<int> source)
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
    public delegate void Notifier8702(int value);

    public sealed class Source8702
    {
        public event Notifier8702 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8703(int value);

    public sealed class Source8703
    {
        public event Notifier8703 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale8704<T>(T value) where T : struct
    {
        int number = 91;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public delegate void Notifier8705(int value);

    public sealed class Source8705
    {
        public event Notifier8705 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
