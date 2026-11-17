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

    public interface IHandler9230
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9230(string Label, int Payload);

    public sealed class Handler9230 : IHandler9230
    {
        public void Handle(int value) { }
    }
    public struct Point9231
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid9231
    {
        public static Point9231 Make(int left, int right)
        {
            return new Point9231 { X = left, Y = right };
        }
    }
    public static int Scale9232<T>(T value) where T : struct
    {
        int number = 87;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
    public static int Scale9233<T>(T value) where T : struct
    {
        int number = 94;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
    public static int Scale9234<T>(T value) where T : struct
    {
        int number = 101;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 48; }
        return number;
    }
    public sealed class SlotMap9235
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
