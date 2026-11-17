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

    public delegate void Notifier1430(int value);

    public sealed class Source1430
    {
        public event Notifier1430 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap1431
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
    public struct Point1432
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid1432
    {
        public static Point1432 Make(int left, int right)
        {
            return new Point1432 { X = left, Y = right };
        }
    }
    public struct Point1433
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 42;
    }

    public static class Grid1433
    {
        public static Point1433 Make(int left, int right)
        {
            return new Point1433 { X = left, Y = right };
        }
    }
    public interface IHandler1434
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1434(string Label, int Payload);

    public sealed class Handler1434 : IHandler1434
    {
        public void Handle(int value) { }
    }
    public interface IHandler1435
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1435(string Label, int Payload);

    public sealed class Handler1435 : IHandler1435
    {
        public void Handle(int value) { }
    }
}
