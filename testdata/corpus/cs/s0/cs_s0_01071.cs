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

    public interface IHandler10030
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10030(string Label, int Payload);

    public sealed class Handler10030 : IHandler10030
    {
        public void Handle(int value) { }
    }
    public struct Point10031
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid10031
    {
        public static Point10031 Make(int left, int right)
        {
            return new Point10031 { X = left, Y = right };
        }
    }
    public struct Point10032
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid10032
    {
        public static Point10032 Make(int left, int right)
        {
            return new Point10032 { X = left, Y = right };
        }
    }
    public sealed class SlotMap10033
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
