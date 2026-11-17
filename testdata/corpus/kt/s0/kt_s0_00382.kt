package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3150 {
    data class Success(val payload: String) : Outcome3150()
    data object Failure : Outcome3150()
}

fun describe3150(outcome: Outcome3150): String = when (outcome) {
    is Outcome3150.Success -> outcome.payload
    Outcome3150.Failure -> "failed"
}
data class Pair3151<A, B>(val first: A, val second: B)

fun combine3151(left: Int, right: Int): Pair3151<Int, Int> {
    return Pair3151(left + 33, right + 29)
}
@JvmName("step3152")
fun step3152(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
data class Pair3153<A, B>(val first: A, val second: B)

fun combine3153(left: Int, right: Int): Pair3153<Int, Int> {
    return Pair3153(left + 47, right + 51)
}
@JvmName("step3154")
fun step3154(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
class Counter3155 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
