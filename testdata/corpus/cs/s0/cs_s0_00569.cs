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

    public sealed class SlotMap5330
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
    public static int Scale5331<T>(T value) where T : struct
    {
        int number = 54;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public struct Point5332
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid5332
    {
        public static Point5332 Make(int left, int right)
        {
            return new Point5332 { X = left, Y = right };
        }
    }
    public delegate void Notifier5333(int value);

    public sealed class Source5333
    {
        public event Notifier5333 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5334<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 63; }
        return number;
    }
    public interface IHandler5335
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5335(string Label, int Payload);

    public sealed class Handler5335 : IHandler5335
    {
        public void Handle(int value) { }
    }
}
