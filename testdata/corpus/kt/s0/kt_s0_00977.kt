package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8010 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
data class Pair8011<A, B>(val first: A, val second: B)

fun combine8011(left: Int, right: Int): Pair8011<Int, Int> {
    return Pair8011(left + 46, right + 15)
}
@JvmName("step8012")
fun step8012(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
data class Pair8013<A, B>(val first: A, val second: B)

fun combine8013(left: Int, right: Int): Pair8013<Int, Int> {
    return Pair8013(left + 60, right + 37)
}
sealed class Outcome8014 {
    data class Success(val payload: String) : Outcome8014()
    data object Failure : Outcome8014()
}

fun describe8014(outcome: Outcome8014): String = when (outcome) {
    is Outcome8014.Success -> outcome.payload
    Outcome8014.Failure -> "failed"
}
