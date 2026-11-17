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

    public interface IHandler6270
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6270(string Label, int Payload);

    public sealed class Handler6270 : IHandler6270
    {
        public void Handle(int value) { }
    }
    public struct Point6271
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid6271
    {
        public static Point6271 Make(int left, int right)
        {
            return new Point6271 { X = left, Y = right };
        }
    }
    public struct Point6272
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 22;
    }

    public static class Grid6272
    {
        public static Point6272 Make(int left, int right)
        {
            return new Point6272 { X = left, Y = right };
        }
    }
    public sealed class SlotMap6273
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
    public struct Point6274
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid6274
    {
        public static Point6274 Make(int left, int right)
        {
            return new Point6274 { X = left, Y = right };
        }
    }
}
