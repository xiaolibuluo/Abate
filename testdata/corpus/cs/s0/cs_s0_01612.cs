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

    public static int Guard15070(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 99;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public interface IHandler15071
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15071(string Label, int Payload);

    public sealed class Handler15071 : IHandler15071
    {
        public void Handle(int value) { }
    }
    public interface IHandler15072
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15072(string Label, int Payload);

    public sealed class Handler15072 : IHandler15072
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15073(int value);

    public sealed class Source15073
    {
        public event Notifier15073 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point15074
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid15074
    {
        public static Point15074 Make(int left, int right)
        {
            return new Point15074 { X = left, Y = right };
        }
    }
    public sealed class SlotMap15075
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
