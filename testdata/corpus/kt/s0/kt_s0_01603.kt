package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13230<A, B>(val first: A, val second: B)

fun combine13230(left: Int, right: Int): Pair13230<Int, Int> {
    return Pair13230(left + 35, right + 59)
}
@JvmName("step13231")
fun step13231(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
sealed class Outcome13232 {
    data class Success(val payload: String) : Outcome13232()
    data object Failure : Outcome13232()
}

fun describe13232(outcome: Outcome13232): String = when (outcome) {
    is Outcome13232.Success -> outcome.payload
    Outcome13232.Failure -> "failed"
}
