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

    public delegate void Notifier16020(int value);

    public sealed class Source16020
    {
        public event Notifier16020 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap16021
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
    public static int Scale16022<T>(T value) where T : struct
    {
        int number = 87;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public static int Scale16023<T>(T value) where T : struct
    {
        int number = 94;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier16024(int value);

    public sealed class Source16024
    {
        public event Notifier16024 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard16025(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
