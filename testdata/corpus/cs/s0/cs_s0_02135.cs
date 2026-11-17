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

    public sealed class SlotMap19940
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
    public static int Scale19941<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public delegate void Notifier19942(int value);

    public sealed class Source19942
    {
        public event Notifier19942 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19943(int value);

    public sealed class Source19943
    {
        public event Notifier19943 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19944(int value);

    public sealed class Source19944
    {
        public event Notifier19944 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19945(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 63;
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
    public delegate void Notifier19946(int value);

    public sealed class Source19946
    {
        public event Notifier19946 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
