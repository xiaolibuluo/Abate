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

    public delegate void Notifier5300(int value);

    public sealed class Source5300
    {
        public event Notifier5300 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5301<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public sealed class SlotMap5302
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
    public interface IHandler5303
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5303(string Label, int Payload);

    public sealed class Handler5303 : IHandler5303
    {
        public void Handle(int value) { }
    }
    public struct Point5304
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid5304
    {
        public static Point5304 Make(int left, int right)
        {
            return new Point5304 { X = left, Y = right };
        }
    }
    public struct Point5305
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid5305
    {
        public static Point5305 Make(int left, int right)
        {
            return new Point5305 { X = left, Y = right };
        }
    }
}
