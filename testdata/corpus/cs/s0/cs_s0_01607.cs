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

    public delegate void Notifier15030(int value);

    public sealed class Source15030
    {
        public event Notifier15030 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale15031<T>(T value) where T : struct
    {
        int number = 54;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public struct Point15032
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid15032
    {
        public static Point15032 Make(int left, int right)
        {
            return new Point15032 { X = left, Y = right };
        }
    }
    public sealed class SlotMap15033
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
    public delegate void Notifier15034(int value);

    public sealed class Source15034
    {
        public event Notifier15034 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler15035
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15035(string Label, int Payload);

    public sealed class Handler15035 : IHandler15035
    {
        public void Handle(int value) { }
    }
}
