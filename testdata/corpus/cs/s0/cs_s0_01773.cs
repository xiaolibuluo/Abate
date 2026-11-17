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

    public interface IHandler16530
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16530(string Label, int Payload);

    public sealed class Handler16530 : IHandler16530
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate16531(IEnumerable<int> source)
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
    public delegate void Notifier16532(int value);

    public sealed class Source16532
    {
        public event Notifier16532 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale16533<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public delegate void Notifier16534(int value);

    public sealed class Source16534
    {
        public event Notifier16534 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler16535
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16535(string Label, int Payload);

    public sealed class Handler16535 : IHandler16535
    {
        public void Handle(int value) { }
    }
    public interface IHandler16536
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16536(string Label, int Payload);

    public sealed class Handler16536 : IHandler16536
    {
        public void Handle(int value) { }
    }
}
