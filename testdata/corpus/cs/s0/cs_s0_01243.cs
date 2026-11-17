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

    public static int Guard11650(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 18;
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
    public static int Guard11651(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
    public async Task<int> Aggregate11652(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public delegate void Notifier11653(int value);

    public sealed class Source11653
    {
        public event Notifier11653 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler11654
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11654(string Label, int Payload);

    public sealed class Handler11654 : IHandler11654
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11655(int value);

    public sealed class Source11655
    {
        public event Notifier11655 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
