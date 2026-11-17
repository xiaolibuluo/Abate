package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8610 {
    data class Success(val payload: String) : Outcome8610()
    data object Failure : Outcome8610()
}

fun describe8610(outcome: Outcome8610): String = when (outcome) {
    is Outcome8610.Success -> outcome.payload
    Outcome8610.Failure -> "failed"
}
data class Pair8611<A, B>(val first: A, val second: B)

fun combine8611(left: Int, right: Int): Pair8611<Int, Int> {
    return Pair8611(left + 50, right + 69)
}
sealed class Outcome8612 {
    data class Success(val payload: String) : Outcome8612()
    data object Failure : Outcome8612()
}

fun describe8612(outcome: Outcome8612): String = when (outcome) {
    is Outcome8612.Success -> outcome.payload
    Outcome8612.Failure -> "failed"
}
@JvmName("step8613")
fun step8613(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
sealed class Outcome8614 {
    data class Success(val payload: String) : Outcome8614()
    data object Failure : Outcome8614()
}

fun describe8614(outcome: Outcome8614): String = when (outcome) {
    is Outcome8614.Success -> outcome.payload
    Outcome8614.Failure -> "failed"
}
@JvmName("step8615")
fun step8615(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
@JvmName("step8616")
fun step8616(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
