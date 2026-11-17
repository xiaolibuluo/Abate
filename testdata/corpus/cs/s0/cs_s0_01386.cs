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

    public struct Point12980
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid12980
    {
        public static Point12980 Make(int left, int right)
        {
            return new Point12980 { X = left, Y = right };
        }
    }
    public interface IHandler12981
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12981(string Label, int Payload);

    public sealed class Handler12981 : IHandler12981
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap12982
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
    public interface IHandler12983
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12983(string Label, int Payload);

    public sealed class Handler12983 : IHandler12983
    {
        public void Handle(int value) { }
    }
}
