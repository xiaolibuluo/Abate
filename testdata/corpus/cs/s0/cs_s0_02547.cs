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

    public static int Scale23800<T>(T value) where T : struct
    {
        int number = 99;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public delegate void Notifier23801(int value);

    public sealed class Source23801
    {
        public event Notifier23801 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23802
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid23802
    {
        public static Point23802 Make(int left, int right)
        {
            return new Point23802 { X = left, Y = right };
        }
    }
    public sealed class SlotMap23803
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
    public struct Point23804
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid23804
    {
        public static Point23804 Make(int left, int right)
        {
            return new Point23804 { X = left, Y = right };
        }
    }
    public interface IHandler23805
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23805(string Label, int Payload);

    public sealed class Handler23805 : IHandler23805
    {
        public void Handle(int value) { }
    }
}
