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

    public delegate void Notifier21820(int value);

    public sealed class Source21820
    {
        public event Notifier21820 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap21821
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
    public static int Scale21822<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 55; }
        return number;
    }
    public struct Point21823
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 68;
    }

    public static class Grid21823
    {
        public static Point21823 Make(int left, int right)
        {
            return new Point21823 { X = left, Y = right };
        }
    }
    public delegate void Notifier21824(int value);

    public sealed class Source21824
    {
        public event Notifier21824 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21825
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21825(string Label, int Payload);

    public sealed class Handler21825 : IHandler21825
    {
        public void Handle(int value) { }
    }
}
