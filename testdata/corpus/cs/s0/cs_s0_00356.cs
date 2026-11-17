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

    public delegate void Notifier3370(int value);

    public sealed class Source3370
    {
        public event Notifier3370 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier3371(int value);

    public sealed class Source3371
    {
        public event Notifier3371 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate3372(IEnumerable<int> source)
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
    public interface IHandler3373
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3373(string Label, int Payload);

    public sealed class Handler3373 : IHandler3373
    {
        public void Handle(int value) { }
    }
    public interface IHandler3374
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3374(string Label, int Payload);

    public sealed class Handler3374 : IHandler3374
    {
        public void Handle(int value) { }
    }
    public static int Guard3375(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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
