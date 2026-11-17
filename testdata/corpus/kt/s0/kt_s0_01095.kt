package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8980 {
    data class Success(val payload: String) : Outcome8980()
    data object Failure : Outcome8980()
}

fun describe8980(outcome: Outcome8980): String = when (outcome) {
    is Outcome8980.Success -> outcome.payload
    Outcome8980.Failure -> "failed"
}
@JvmName("step8981")
fun step8981(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
sealed class Outcome8982 {
    data class Success(val payload: String) : Outcome8982()
    data object Failure : Outcome8982()
}

fun describe8982(outcome: Outcome8982): String = when (outcome) {
    is Outcome8982.Success -> outcome.payload
    Outcome8982.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8983(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8984 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8984 = "registry-67"
    }
}
