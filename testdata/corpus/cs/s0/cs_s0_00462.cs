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

    public static int Guard4330(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 8;
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
    public delegate void Notifier4331(int value);

    public sealed class Source4331
    {
        public event Notifier4331 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4332(int value);

    public sealed class Source4332
    {
        public event Notifier4332 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4333(int value);

    public sealed class Source4333
    {
        public event Notifier4333 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4334(int value);

    public sealed class Source4334
    {
        public event Notifier4334 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap4335
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
