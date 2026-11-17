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

    public delegate void Notifier8260(int value);

    public sealed class Source8260
    {
        public event Notifier8260 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8261(int value);

    public sealed class Source8261
    {
        public event Notifier8261 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap8262
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
    public sealed class SlotMap8263
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
    public struct Point8264
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid8264
    {
        public static Point8264 Make(int left, int right)
        {
            return new Point8264 { X = left, Y = right };
        }
    }
}
