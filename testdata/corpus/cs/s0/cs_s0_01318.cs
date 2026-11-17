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

    public delegate void Notifier12340(int value);

    public sealed class Source12340
    {
        public event Notifier12340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point12341
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 49;
    }

    public static class Grid12341
    {
        public static Point12341 Make(int left, int right)
        {
            return new Point12341 { X = left, Y = right };
        }
    }
    public sealed class SlotMap12342
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
