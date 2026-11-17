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

    public static int Scale10190<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
    public interface IHandler10191
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10191(string Label, int Payload);

    public sealed class Handler10191 : IHandler10191
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier10192(int value);

    public sealed class Source10192
    {
        public event Notifier10192 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap10193
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
    public struct Point10194
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid10194
    {
        public static Point10194 Make(int left, int right)
        {
            return new Point10194 { X = left, Y = right };
        }
    }
    public static int Scale10195<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
}
