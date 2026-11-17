package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18190")
fun step18190(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
sealed class Outcome18191 {
    data class Success(val payload: String) : Outcome18191()
    data object Failure : Outcome18191()
}

fun describe18191(outcome: Outcome18191): String = when (outcome) {
    is Outcome18191.Success -> outcome.payload
    Outcome18191.Failure -> "failed"
}
@JvmName("step18192")
fun step18192(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
@JvmName("step18193")
fun step18193(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
object Registry18194 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18194 = "registry-12"
    }
}
