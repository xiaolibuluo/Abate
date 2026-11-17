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

    public sealed class SlotMap16980
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
    public delegate void Notifier16981(int value);

    public sealed class Source16981
    {
        public event Notifier16981 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16982
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid16982
    {
        public static Point16982 Make(int left, int right)
        {
            return new Point16982 { X = left, Y = right };
        }
    }
    public delegate void Notifier16983(int value);

    public sealed class Source16983
    {
        public event Notifier16983 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale16984<T>(T value) where T : struct
    {
        int number = 88;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public interface IHandler16985
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16985(string Label, int Payload);

    public sealed class Handler16985 : IHandler16985
    {
        public void Handle(int value) { }
    }
}
