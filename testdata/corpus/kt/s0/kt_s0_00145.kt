package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1170")
fun step1170(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
fun <T, R> List<T>.mapChecked1171(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome1172 {
    data class Success(val payload: String) : Outcome1172()
    data object Failure : Outcome1172()
}

fun describe1172(outcome: Outcome1172): String = when (outcome) {
    is Outcome1172.Success -> outcome.payload
    Outcome1172.Failure -> "failed"
}
@JvmName("step1173")
fun step1173(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
object Registry1174 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1174 = "registry-99"
    }
}
