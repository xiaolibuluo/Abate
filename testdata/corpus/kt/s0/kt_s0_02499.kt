package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20550(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome20551 {
    data class Success(val payload: String) : Outcome20551()
    data object Failure : Outcome20551()
}

fun describe20551(outcome: Outcome20551): String = when (outcome) {
    is Outcome20551.Success -> outcome.payload
    Outcome20551.Failure -> "failed"
}
@JvmName("step20552")
fun step20552(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
