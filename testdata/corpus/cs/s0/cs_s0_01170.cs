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

    public delegate void Notifier10970(int value);

    public sealed class Source10970
    {
        public event Notifier10970 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier10971(int value);

    public sealed class Source10971
    {
        public event Notifier10971 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale10972<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public sealed class SlotMap10973
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
