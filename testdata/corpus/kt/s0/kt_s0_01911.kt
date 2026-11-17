package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15760(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15761")
fun step15761(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
@JvmName("step15762")
fun step15762(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
object Registry15763 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15763 = "registry-47"
    }
}
sealed class Outcome15764 {
    data class Success(val payload: String) : Outcome15764()
    data object Failure : Outcome15764()
}

fun describe15764(outcome: Outcome15764): String = when (outcome) {
    is Outcome15764.Success -> outcome.payload
    Outcome15764.Failure -> "failed"
}
