package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3750")
fun step3750(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair3751<A, B>(val first: A, val second: B)

fun combine3751(left: Int, right: Int): Pair3751<Int, Int> {
    return Pair3751(left + 37, right + 22)
}
@JvmName("step3752")
fun step3752(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
sealed class Outcome3753 {
    data class Success(val payload: String) : Outcome3753()
    data object Failure : Outcome3753()
}

fun describe3753(outcome: Outcome3753): String = when (outcome) {
    is Outcome3753.Success -> outcome.payload
    Outcome3753.Failure -> "failed"
}
class Counter3754 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
class Counter3755 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
