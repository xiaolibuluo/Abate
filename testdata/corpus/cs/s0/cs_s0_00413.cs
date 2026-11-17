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

    public interface IHandler3880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3880(string Label, int Payload);

    public sealed class Handler3880 : IHandler3880
    {
        public void Handle(int value) { }
    }
    public struct Point3881
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 12;
    }

    public static class Grid3881
    {
        public static Point3881 Make(int left, int right)
        {
            return new Point3881 { X = left, Y = right };
        }
    }
    public sealed class SlotMap3882
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
    public interface IHandler3883
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3883(string Label, int Payload);

    public sealed class Handler3883 : IHandler3883
    {
        public void Handle(int value) { }
    }
    public struct Point3884
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid3884
    {
        public static Point3884 Make(int left, int right)
        {
            return new Point3884 { X = left, Y = right };
        }
    }
    public interface IHandler3885
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3885(string Label, int Payload);

    public sealed class Handler3885 : IHandler3885
    {
        public void Handle(int value) { }
    }
}
