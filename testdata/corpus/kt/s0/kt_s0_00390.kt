package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3210 {
    data class Success(val payload: String) : Outcome3210()
    data object Failure : Outcome3210()
}

fun describe3210(outcome: Outcome3210): String = when (outcome) {
    is Outcome3210.Success -> outcome.payload
    Outcome3210.Failure -> "failed"
}
class Counter3211 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
@JvmName("step3212")
fun step3212(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
data class Pair3213<A, B>(val first: A, val second: B)

fun combine3213(left: Int, right: Int): Pair3213<Int, Int> {
    return Pair3213(left + 28, right + 32)
}
class Counter3214 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
