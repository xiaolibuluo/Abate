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

    public delegate void Notifier22820(int value);

    public sealed class Source22820
    {
        public event Notifier22820 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard22821(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 93;
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
    public sealed class SlotMap22822
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
    public interface IHandler22823
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22823(string Label, int Payload);

    public sealed class Handler22823 : IHandler22823
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22824(int value);

    public sealed class Source22824
    {
        public event Notifier22824 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler22825
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22825(string Label, int Payload);

    public sealed class Handler22825 : IHandler22825
    {
        public void Handle(int value) { }
    }
}
