package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15750 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15750 = "registry-13"
    }
}
@JvmName("step15751")
fun step15751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
sealed class Outcome15752 {
    data class Success(val payload: String) : Outcome15752()
    data object Failure : Outcome15752()
}

fun describe15752(outcome: Outcome15752): String = when (outcome) {
    is Outcome15752.Success -> outcome.payload
    Outcome15752.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15753(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
