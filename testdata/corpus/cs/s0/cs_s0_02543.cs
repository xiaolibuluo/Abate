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

    public delegate void Notifier23760(int value);

    public sealed class Source23760
    {
        public event Notifier23760 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier23761(int value);

    public sealed class Source23761
    {
        public event Notifier23761 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23762<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public static int Guard23763(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
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
    public sealed class SlotMap23764
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
