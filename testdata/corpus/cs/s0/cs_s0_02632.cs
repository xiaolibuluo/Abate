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

    public delegate void Notifier24620(int value);

    public sealed class Source24620
    {
        public event Notifier24620 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler24621
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24621(string Label, int Payload);

    public sealed class Handler24621 : IHandler24621
    {
        public void Handle(int value) { }
    }
    public static int Guard24622(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 15;
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
    public async Task<int> Aggregate24623(IEnumerable<int> source)
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
