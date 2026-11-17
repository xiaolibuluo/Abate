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

    public static int Guard20390(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 31;
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
    public delegate void Notifier20391(int value);

    public sealed class Source20391
    {
        public event Notifier20391 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate20392(IEnumerable<int> source)
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
    public interface IHandler20393
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20393(string Label, int Payload);

    public sealed class Handler20393 : IHandler20393
    {
        public void Handle(int value) { }
    }
    public interface IHandler20394
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20394(string Label, int Payload);

    public sealed class Handler20394 : IHandler20394
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier20395(int value);

    public sealed class Source20395
    {
        public event Notifier20395 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
