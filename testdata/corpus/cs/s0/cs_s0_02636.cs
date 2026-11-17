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

    public static int Guard24660(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 53;
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
    public delegate void Notifier24661(int value);

    public sealed class Source24661
    {
        public event Notifier24661 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler24662
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24662(string Label, int Payload);

    public sealed class Handler24662 : IHandler24662
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate24663(IEnumerable<int> source)
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
}
