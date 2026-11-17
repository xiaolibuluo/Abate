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

    public struct Point3240
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid3240
    {
        public static Point3240 Make(int left, int right)
        {
            return new Point3240 { X = left, Y = right };
        }
    }
    public interface IHandler3241
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3241(string Label, int Payload);

    public sealed class Handler3241 : IHandler3241
    {
        public void Handle(int value) { }
    }
    public static int Scale3242<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public delegate void Notifier3243(int value);

    public sealed class Source3243
    {
        public event Notifier3243 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap3244
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
