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

    public sealed class SlotMap13110
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
    public struct Point13111
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid13111
    {
        public static Point13111 Make(int left, int right)
        {
            return new Point13111 { X = left, Y = right };
        }
    }
    public sealed class SlotMap13112
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
    public interface IHandler13113
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13113(string Label, int Payload);

    public sealed class Handler13113 : IHandler13113
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13114(int value);

    public sealed class Source13114
    {
        public event Notifier13114 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
