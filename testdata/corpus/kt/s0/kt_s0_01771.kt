package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14530<A, B>(val first: A, val second: B)

fun combine14530(left: Int, right: Int): Pair14530<Int, Int> {
    return Pair14530(left + 76, right + 54)
}
@JvmName("step14531")
fun step14531(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
@JvmName("step14532")
fun step14532(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
@JvmName("step14533")
fun step14533(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
sealed class Outcome14534 {
    data class Success(val payload: String) : Outcome14534()
    data object Failure : Outcome14534()
}

fun describe14534(outcome: Outcome14534): String = when (outcome) {
    is Outcome14534.Success -> outcome.payload
    Outcome14534.Failure -> "failed"
}
sealed class Outcome14535 {
    data class Success(val payload: String) : Outcome14535()
    data object Failure : Outcome14535()
}

fun describe14535(outcome: Outcome14535): String = when (outcome) {
    is Outcome14535.Success -> outcome.payload
    Outcome14535.Failure -> "failed"
}
class Counter14536 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
class Counter14537 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
