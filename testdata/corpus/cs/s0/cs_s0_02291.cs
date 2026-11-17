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

    public sealed class SlotMap21380
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
    public static int Scale21381<T>(T value) where T : struct
    {
        int number = 64;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public delegate void Notifier21382(int value);

    public sealed class Source21382
    {
        public event Notifier21382 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point21383
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid21383
    {
        public static Point21383 Make(int left, int right)
        {
            return new Point21383 { X = left, Y = right };
        }
    }
    public interface IHandler21384
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21384(string Label, int Payload);

    public sealed class Handler21384 : IHandler21384
    {
        public void Handle(int value) { }
    }
    public static int Scale21385<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
}
