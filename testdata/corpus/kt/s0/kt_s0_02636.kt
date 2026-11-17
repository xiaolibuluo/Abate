package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21770 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
sealed class Outcome21771 {
    data class Success(val payload: String) : Outcome21771()
    data object Failure : Outcome21771()
}

fun describe21771(outcome: Outcome21771): String = when (outcome) {
    is Outcome21771.Success -> outcome.payload
    Outcome21771.Failure -> "failed"
}
data class Pair21772<A, B>(val first: A, val second: B)

fun combine21772(left: Int, right: Int): Pair21772<Int, Int> {
    return Pair21772(left + 93, right + 20)
}
data class Pair21773<A, B>(val first: A, val second: B)

fun combine21773(left: Int, right: Int): Pair21773<Int, Int> {
    return Pair21773(left + 100, right + 31)
}
class Counter21774 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
sealed class Outcome21775 {
    data class Success(val payload: String) : Outcome21775()
    data object Failure : Outcome21775()
}

fun describe21775(outcome: Outcome21775): String = when (outcome) {
    is Outcome21775.Success -> outcome.payload
    Outcome21775.Failure -> "failed"
}
