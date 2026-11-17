package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14350 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14350 = "registry-36"
    }
}
object Registry14351 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14351 = "registry-43"
    }
}
fun <T, R> List<T>.mapChecked14352(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome14353 {
    data class Success(val payload: String) : Outcome14353()
    data object Failure : Outcome14353()
}

fun describe14353(outcome: Outcome14353): String = when (outcome) {
    is Outcome14353.Success -> outcome.payload
    Outcome14353.Failure -> "failed"
}
@JvmName("step14354")
fun step14354(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
