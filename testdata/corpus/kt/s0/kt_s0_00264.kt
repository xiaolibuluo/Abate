package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2160 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2160 = "registry-97"
    }
}
sealed class Outcome2161 {
    data class Success(val payload: String) : Outcome2161()
    data object Failure : Outcome2161()
}

fun describe2161(outcome: Outcome2161): String = when (outcome) {
    is Outcome2161.Success -> outcome.payload
    Outcome2161.Failure -> "failed"
}
object Registry2162 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2162 = "registry-14"
    }
}
fun <T, R> List<T>.mapChecked2163(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry2164 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2164 = "registry-28"
    }
}
