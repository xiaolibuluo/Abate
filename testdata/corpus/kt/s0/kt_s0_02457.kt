package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20190 {
    data class Success(val payload: String) : Outcome20190()
    data object Failure : Outcome20190()
}

fun describe20190(outcome: Outcome20190): String = when (outcome) {
    is Outcome20190.Success -> outcome.payload
    Outcome20190.Failure -> "failed"
}
@JvmName("step20191")
fun step20191(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
data class Pair20192<A, B>(val first: A, val second: B)

fun combine20192(left: Int, right: Int): Pair20192<Int, Int> {
    return Pair20192(left + 76, right + 12)
}
sealed class Outcome20193 {
    data class Success(val payload: String) : Outcome20193()
    data object Failure : Outcome20193()
}

fun describe20193(outcome: Outcome20193): String = when (outcome) {
    is Outcome20193.Success -> outcome.payload
    Outcome20193.Failure -> "failed"
}
class Counter20194 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
