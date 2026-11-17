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

    public struct Point11500
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 17;
    }

    public static class Grid11500
    {
        public static Point11500 Make(int left, int right)
        {
            return new Point11500 { X = left, Y = right };
        }
    }
    public struct Point11501
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 24;
    }

    public static class Grid11501
    {
        public static Point11501 Make(int left, int right)
        {
            return new Point11501 { X = left, Y = right };
        }
    }
    public interface IHandler11502
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11502(string Label, int Payload);

    public sealed class Handler11502 : IHandler11502
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap11503
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
