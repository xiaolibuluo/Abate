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

    public interface IHandler23830
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23830(string Label, int Payload);

    public sealed class Handler23830 : IHandler23830
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate23831(IEnumerable<int> source)
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
    public delegate void Notifier23832(int value);

    public sealed class Source23832
    {
        public event Notifier23832 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard23833(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 62;
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
    public delegate void Notifier23834(int value);

    public sealed class Source23834
    {
        public event Notifier23834 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23835
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23835(string Label, int Payload);

    public sealed class Handler23835 : IHandler23835
    {
        public void Handle(int value) { }
    }
}
