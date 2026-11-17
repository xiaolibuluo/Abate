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

    public delegate void Notifier6810(int value);

    public sealed class Source6810
    {
        public event Notifier6810 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6811(int value);

    public sealed class Source6811
    {
        public event Notifier6811 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6812(int value);

    public sealed class Source6812
    {
        public event Notifier6812 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6813
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6813(string Label, int Payload);

    public sealed class Handler6813 : IHandler6813
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier6814(int value);

    public sealed class Source6814
    {
        public event Notifier6814 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6815
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6815(string Label, int Payload);

    public sealed class Handler6815 : IHandler6815
    {
        public void Handle(int value) { }
    }
    public static int Guard6816(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 73;
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
    public sealed class SlotMap6817
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
