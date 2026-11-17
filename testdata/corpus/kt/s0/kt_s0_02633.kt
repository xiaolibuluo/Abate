package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21740 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
class Counter21741 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
class Counter21742 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome21743 {
    data class Success(val payload: String) : Outcome21743()
    data object Failure : Outcome21743()
}

fun describe21743(outcome: Outcome21743): String = when (outcome) {
    is Outcome21743.Success -> outcome.payload
    Outcome21743.Failure -> "failed"
}
data class Pair21744<A, B>(val first: A, val second: B)

fun combine21744(left: Int, right: Int): Pair21744<Int, Int> {
    return Pair21744(left + 68, right + 21)
}
