package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17260 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17260 = "registry-36"
    }
}
fun <T, R> List<T>.mapChecked17261(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17262 {
    data class Success(val payload: String) : Outcome17262()
    data object Failure : Outcome17262()
}

fun describe17262(outcome: Outcome17262): String = when (outcome) {
    is Outcome17262.Success -> outcome.payload
    Outcome17262.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17263(): T? = if (size > 1) this[1] else null

fun render17263(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17264")
fun step17264(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
