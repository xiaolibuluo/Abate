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

    public struct Point7730
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid7730
    {
        public static Point7730 Make(int left, int right)
        {
            return new Point7730 { X = left, Y = right };
        }
    }
    public delegate void Notifier7731(int value);

    public sealed class Source7731
    {
        public event Notifier7731 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier7732(int value);

    public sealed class Source7732
    {
        public event Notifier7732 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point7733
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 84;
    }

    public static class Grid7733
    {
        public static Point7733 Make(int left, int right)
        {
            return new Point7733 { X = left, Y = right };
        }
    }
    public sealed class SlotMap7734
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
