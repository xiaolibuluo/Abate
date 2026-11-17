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

    public delegate void Notifier17960(int value);

    public sealed class Source17960
    {
        public event Notifier17960 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap17961
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
    public interface IHandler17962
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17962(string Label, int Payload);

    public sealed class Handler17962 : IHandler17962
    {
        public void Handle(int value) { }
    }
    public interface IHandler17963
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17963(string Label, int Payload);

    public sealed class Handler17963 : IHandler17963
    {
        public void Handle(int value) { }
    }
    public struct Point17964
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid17964
    {
        public static Point17964 Make(int left, int right)
        {
            return new Point17964 { X = left, Y = right };
        }
    }
    public static int Guard17965(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
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
}
