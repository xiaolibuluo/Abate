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

    public static int Scale6720<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public static int Scale6721<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public delegate void Notifier6722(int value);

    public sealed class Source6722
    {
        public event Notifier6722 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap6723
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
    public sealed class SlotMap6724
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
